.PHONY: gen-user gen-product gen-cart

gen-user:
	cwgo server --type RPC --module zqzqsb/gomall/app/user -I /home/zq/Projects/GoMall/GomallBackend/idl --idl /home/zq/Projects/GoMall/GomallBackend/idl/user.proto --service user --hex

gen-product:
	cwgo server --type RPC --module zqzqsb/gomall/app/product -I /home/zq/Projects/GoMall/GomallBackend/idl --idl /home/zq/Projects/GoMall/GomallBackend/idl/product.proto --service product --hex

gen-cart:
	cwgo server --type RPC --module zqzqsb/gomall/app/cart -I /home/zq/Projects/GoMall/GomallBackend/idl --idl /home/zq/Projects/GoMall/GomallBackend/idl/cart.proto --service cart --hex

.PHONY: gen-all
gen-all: gen-user gen-product gen-cart