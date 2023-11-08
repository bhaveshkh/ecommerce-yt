package database

import (
	"errors"
)

var (
	ErrCantFindProduct    = errors.New("cant find the product")
	ErrCantDecodeProducts = errors.New("cant fint he product")
	ErrUserIdIsNotValid   = errors.New("this user is not valid")
	ErrCantRemoveItemCart = errors.New("cannot add this product to cart")
	ErrCantUpdateUser     = errors.New("cannot remove this item from the cart")
	ErrCantGetItem        = errors.New("was unable to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("cannot update the purchase")
)

func AddProductToCart() {}

func RemoveCartItem() {}

func BuytItemFromCart() {}

func InstantBuyer() {}
