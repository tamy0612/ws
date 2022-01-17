package dict

// Item is a container for each dictionary item
type Item struct {
    ID uint32       `db:"item_id"`
    Word string     `db:"word"`
    Mean string     `db:"mean"`
    Level uint32    `db:"level"`
}
