//By Ahmad Fadhil Rizqi

package models

import "gorm.io/gorm"

// Transaction mewakili transaksi pembelian barang
type Transaction struct {
    gorm.Model
    ItemID  uint   `json:"item_id"`
    BuyerID uint   `json:"buyer_id"`
    Status  string `json:"status"`
    Date    string `json:"date"` 
}
