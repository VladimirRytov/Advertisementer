#!/bin/bash

trap "exit 1" ERR

echo "========> building sqlite for testing <========"
go install github.com/mattn/go-sqlite3

echo "========> Testing advertisements package <========"
go test github.com/VladimirRytov/advertisementer/internal/advertisements
echo "========> Testing logging package <========"
go test github.com/VladimirRytov/advertisementer/internal/logging
echo "========> Testing mapper package <========"
go test github.com/VladimirRytov/advertisementer/internal/mapper
echo "========> Testing systeminfo package <========"
go test github.com/VladimirRytov/advertisementer/internal/systeminfo
echo "========> Testing encodedecoder package <========"
go test github.com/VladimirRytov/advertisementer/internal/encodedecoder
echo "========> Testing encryptor package <========"
go test github.com/VladimirRytov/advertisementer/internal/encryptor
echo "========> Testing filestorage package <========"
go test github.com/VladimirRytov/advertisementer/internal/filestorage

echo "========> Testing configstorage package <========"
go test github.com/VladimirRytov/advertisementer/internal/handlers/confighandler
echo "========> Testing advreport package <========"
go test github.com/VladimirRytov/advertisementer/internal/handlers/advreport
echo "========> Testing exelexporter package <========"
go test github.com/VladimirRytov/advertisementer/internal/handlers/advreport/exelexporter
echo "========> Testing costRateCalculatorHandler package <========"
go test github.com/VladimirRytov/advertisementer/internal/handlers/costcalculationhandler
# echo "========> Testing exportHandler package <========"
# go test
# echo "========> Testing importHandler package <========"
# go test

trap "rm internal/datastorage/sql/orm/testing.sqlite; exit 1" ERR

echo "========> Testing orm package <========"
echo "========> Testing tools <========"
go test -run '^(TestFetchTagsName|TestFetchExtraChargeName|TestReleaseDates)$' github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm
echo "========> Testing mapper <========"
go test -run '^TestConvert' github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm
go test -run 'ToModel$' github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm
go test -run 'ToDto$' github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm
echo "========> Testing connection <========"
go test -run '^TestConnectTo' github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm
echo "========> Testing create requests <========"
go test -timeout 30s -run '^TestCreate' github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm
echo "========> Testing get requests <========"
go test -run '^(TestClientByID|TestOrdersByID|TestLineAdvertisementByID|TestBlockAdvertisementByID|TestTagByName|TestExtraChargeByName|TestCostRateByName)$' github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm
echo "========> Testing search requests <========"
go test -timeout 30s -run \
  '^(TestAllClients|TestAllOrders|TestOrdersByClientName|TestAllLineAdvertisements|TestLineAdvertisementsByOrderID|TestAllBlockAdvertisements|TestBlockAdvertisementsByOrderID|TestBlockAdvertisementBetweenReleaseDates|TestBlockAdvertisementFromReleaseDates|TestLineAdvertisementBetweenReleaseDates|TestLineAdvertisementFromReleaseDates|TestAllTags|TestAllExtraChargess|TestAllCostRates)$' \
  github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm
echo "========> Testing update requests <========"
go test -timeout 30s -run '^TestUpdate' github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm
echo "========> Testing remove requests <========"
go test -run '^TestRemove' github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm
rm internal/datastorage/sql/orm/testing.sqlite