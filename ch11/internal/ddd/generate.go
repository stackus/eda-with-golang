package ddd

//go:generate mockery --name ".*(Aggregate|Entity|Subscriber|Publisher|Handler)$"  --inpackage --case underscore
