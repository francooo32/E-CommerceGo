package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("Can´t find the product")
	ErrCantDecodeProducts = errors.New("Can´t decode the product")
	ErrUserIdIsNotValid   = errors.New("User id is invalid")
	ErrCantUpdateUser     = errors.New("Can´t update user")
	ErrCantRemoveItemCart = errors.New("Can´t remove item from cart")
	ErrCantGetItem        = errors.New("Can´t get the item")
	ErrCantBuyCartItem    = errors.New("Can´t update the selected item")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
