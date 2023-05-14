package repository

func (r *PartnerRepo) Delete(id int) error {
	err := r.db.Delete(&Partner{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}