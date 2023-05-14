package repository

func (r *PartnerRepo) GetFileId(id int) (string, error) {
	var data Partner

	if err := r.db.Where("id = ?", id).First(&data).Error; err != nil {
		return "", err
	}
	return data.FileId, nil
}