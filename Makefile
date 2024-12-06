.PHONY: gen-user
gen-user:
	cwgo server --type RPC  --module zqzqsb/gomall/app/user -I ../../idl --idl ../../idl/user/user.proto --service user --hex