package mysql

import (
	"context"

	"github.com/strugglehonor/KCS/internal/model"
	"gorm.io/gorm"
)

type Pod struct {
	db *gorm.DB
}

func newPod(ms *mysqlstore) *Pod {
	return &Pod{db: ms.db}
}

func (p Pod) Create(ctx context.Context, pod *model.Pod) error {
	return p.db.Create(pod).Error
}

func (p Pod) Delete(ctx context.Context, pod *model.Pod) error {
	return p.db.Delete(pod).Error
}

func (p Pod) GetPodByName(ctx context.Context, podName string) (*model.Pod, error) {
	var pod model.Pod
	
	if err := p.db.Where("name = ?", podName).Error; err != nil {
		return nil, err
	}
	
	p.db.Where("name = ?", podName).First(&pod)
	return &pod, nil
}

func (p Pod) InsertPod(ctx context.Context, pod *model.Pod) error {
	return p.db.Create(pod).Error
}

func (p Pod) Update(ctx context.Context, pod *model.Pod) error {
	return p.db.Save(pod).Error
}

func (p Pod) GetAllPod(ctx context.Context, limit, offset *int64) ([]*model.Pod, error) {
	pods := []*model.Pod{}
	if err :=  p.db.Limit(int(*limit)).Offset(int(*offset)).Find(pods).Error; err != nil {
		return nil, err
	}

	return pods, nil
}