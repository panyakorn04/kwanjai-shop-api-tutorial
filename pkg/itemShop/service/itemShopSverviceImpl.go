package service

import (
	"github.com/labstack/echo/v4"
	entities "github.com/panyakorn04/kwanjai-shop-api-tutorial/entities"
	_inventoryRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/inventory/repository"
	_itemShopException "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemShop/exception"
	_itemShopModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemShop/model"
	_itemShopRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/itemShop/repository"
	_playerCoinModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/playerCoin/model"
	_playerCoinRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/playerCoin/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository   _itemShopRepository.ItemShopRepository
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository
	inventoryRepository  _inventoryRepository.InventoryRepository
	logger               echo.Logger
}

func NewItemShopServiceImpl(itemShopRepository _itemShopRepository.ItemShopRepository, playerCoinRepository _playerCoinRepository.PlayerCoinRepository,
	inventoryRepository _inventoryRepository.InventoryRepository, logger echo.Logger) ItemShopService {
	return &itemShopServiceImpl{
		itemShopRepository:   itemShopRepository,
		playerCoinRepository: playerCoinRepository,
		inventoryRepository:  inventoryRepository,
		logger:               logger,
	}
}

func (s *itemShopServiceImpl) Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error) {

	itemEntityList, err := s.itemShopRepository.Listing(itemFilter)
	if err != nil {
		return nil, err
	}

	totalItems, err := s.itemShopRepository.Counting(itemFilter)
	if err != nil {
		return nil, err
	}

	size := int64(itemFilter.Paginate.Size)
	page := int64(itemFilter.Paginate.Page)

	totalPage := s.totalPageCalculation(totalItems, size)

	result := s.toItemResultsResponse(itemEntityList, page, totalPage)

	return result, nil
}

func (s *itemShopServiceImpl) Buying(buyingReq *_itemShopModel.BuyingReq) (*_playerCoinModel.PlayerCoin, error) {
	itemEntity, err := s.itemShopRepository.FindByID(buyingReq.ItemID)
	if err != nil {
		return nil, err
	}

	totalPrice := s.totalPriceCalculation(itemEntity.ToItemModel(), buyingReq.Quantity)

	if err := s.playerCoinChecking(buyingReq.PlayerID, totalPrice); err != nil {
		return nil, err
	}

	tx := s.itemShopRepository.TransactionBegin()
	purchaseRecording, err := s.itemShopRepository.PurchaseHistoryRecording(&entities.PurchaseHistory{
		PlayerID:        buyingReq.PlayerID,
		ItemID:          buyingReq.ItemID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		Quantity:        buyingReq.Quantity,
		IsBuying:        true,
	}, tx)

	if err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}
	s.logger.Info("purchase history recorded: %s", purchaseRecording.ID)

	playerCoin, err := s.playerCoinRepository.CoinAdding(&entities.PlayerCoin{
		PlayerID: buyingReq.PlayerID,
		Amount:   -totalPrice,
	}, tx)

	s.logger.Info("player coin updated: %s", playerCoin.Amount)

	inventoryEntity, err := s.inventoryRepository.Filling(tx, buyingReq.PlayerID,
		buyingReq.ItemID,
		int(buyingReq.Quantity),
	)

	s.logger.Info("player inventory updated: %d", len(inventoryEntity))

	if err := s.itemShopRepository.TransactionCommit(tx); err != nil {
		return nil, err
	}

	return playerCoin.ToPlayerCoinMold(), nil
}

func (s *itemShopServiceImpl) Selling(sellingReq *_itemShopModel.SellingReq) (*_playerCoinModel.PlayerCoin, error) {
	itemEntity, err := s.itemShopRepository.FindByID(sellingReq.ItemID)
	if err != nil {
		return nil, err
	}

	totalPrice := s.totalPriceCalculation(itemEntity.ToItemModel(), sellingReq.Quantity)
	totalPrice = totalPrice / 2

	if err := s.playerItemChecking(sellingReq.PlayerID, sellingReq.ItemID, sellingReq.Quantity); err != nil {
		return nil, err
	}

	tx := s.itemShopRepository.TransactionBegin()
	purchaseRecording, err := s.itemShopRepository.PurchaseHistoryRecording(&entities.PurchaseHistory{
		PlayerID:        sellingReq.PlayerID,
		ItemID:          sellingReq.ItemID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		Quantity:        sellingReq.Quantity,
		IsBuying:        false,
	}, tx)

	if err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}
	s.logger.Info("purchase history recorded: %s", purchaseRecording.ID)

	playerCoin, err := s.playerCoinRepository.CoinAdding(&entities.PlayerCoin{
		PlayerID: sellingReq.PlayerID,
		Amount:   totalPrice,
	}, tx)

	s.logger.Info("player coin updated: %s", playerCoin.Amount)

	if err := s.inventoryRepository.Removing(
		sellingReq.PlayerID,
		sellingReq.ItemID,
		int(sellingReq.Quantity),
		tx,
	); err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}
	s.logger.Info("Deleted player item from player's inventory for %d records", sellingReq.Quantity)

	if err := s.itemShopRepository.TransactionCommit(tx); err != nil {
		return nil, err
	}

	return playerCoin.ToPlayerCoinMold(), nil
}

func (s *itemShopServiceImpl) totalPageCalculation(totalItem int64, limit int64) int64 {
	totalPage := totalItem / limit
	if totalItem%limit != 0 {
		totalPage++
	}
	return totalPage
}

func (s *itemShopServiceImpl) toItemResultsResponse(itemEntityList []entities.Item, page, totalPage int64) *_itemShopModel.ItemResult {
	items := make([]*_itemShopModel.Item, 0)

	for _, itemEntity := range itemEntityList {
		items = append(items, itemEntity.ToItemModel())
	}

	return &_itemShopModel.ItemResult{
		Items: items,
		Paginate: _itemShopModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}

func (s *itemShopServiceImpl) totalPriceCalculation(item *_itemShopModel.Item, qty uint) int64 {
	return int64(item.Price) * int64(qty)
}

func (s *itemShopServiceImpl) playerCoinChecking(playerID string, totalPrice int64) error {
	playerCoin, err := s.playerCoinRepository.Showing(playerID)
	if err != nil {
		return err
	}

	if playerCoin.Coin < totalPrice {
		s.logger.Errorf("player coin is not enough")
		return &_itemShopException.CoinNotEnough{}
	}

	return nil
}

func (s *itemShopServiceImpl) playerItemChecking(playerID string, itemID uint64, qty uint) error {
	itemCounting := s.inventoryRepository.PlayerItemCounting(playerID, itemID)

	if int(itemCounting) < int(qty) {
		s.logger.Errorf("player item is not enough")
		return &_itemShopException.ItemNotFound{
			ItemID: itemID,
		}
	}
	return nil
}
