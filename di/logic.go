package di

// NewSublogic サブロジックのインスタンス生成
func NewSublogic(db *gorm.DB) Sublogic {
	return &sublogicst{
		db,
	}
}
