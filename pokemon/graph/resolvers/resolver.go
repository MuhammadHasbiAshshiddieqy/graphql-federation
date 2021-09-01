package resolvers

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import "gorm.io/gorm"

// Resolver return *gorm.DB
type Resolver struct{
  POKEMON *gorm.DB
}
