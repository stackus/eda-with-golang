package am

//go:generate buf generate

//go:generate mockery --name ".*(Subscriber|Publisher|Handler)$"  --inpackage --case underscore
