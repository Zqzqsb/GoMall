.PHONY: gen-frontend
gen-frontend:
	cwgo server --type HTTP --idl ../../idl/frontend/home.proto --module zqzqsb/gomall/app/frontend -I ../../idl --service frontend