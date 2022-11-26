package entity

type AnnouncementImage struct {
	AnnouncementId string `json:"announcement_id"`
	ImageURL       string `json:"image_url"`
}

func NewAnnouncementImage(announcementId string, imageUrl string) (*AnnouncementImage, error) {

	announcement := &AnnouncementImage{
		AnnouncementId: announcementId,
		ImageURL:       imageUrl,
	}

	return announcement, nil
}
