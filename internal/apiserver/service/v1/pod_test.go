package v1

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/strugglehonor/KCS/internal/apiserver/store"
	"github.com/strugglehonor/KCS/internal/apiserver/store/fake"
	"github.com/strugglehonor/KCS/internal/model"
	"github.com/strugglehonor/KCS/pkg/redis"
)

type Suite struct {
	suite.Suite
	mockFactory *store.MockFactory

	mockPodStore         *store.MockPodStore
	mockClusterStore     *store.MockClusterStore
	mockDeplolymentStore *store.MockDeploymentStore
	mockVolumeStore      *store.MockVolumeStore

	pods       []*model.Pod
	cluster    []*model.Cluster
	deployment []*model.Deployment
	volume     []*model.Volume
}

func (s *Suite) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	s.pods = fake.FakePod(10)
	s.deployment = fake.FakeDeployment(10)
	s.cluster = fake.FakeCluster(10)
	s.volume = fake.FakeVolume(10)

	s.mockFactory = store.NewMockFactory(ctrl)

	s.mockClusterStore = store.NewMockClusterStore(ctrl)
	s.mockFactory.EXPECT().Cluster().AnyTimes().Return(s.mockClusterStore)

	s.mockDeplolymentStore = store.NewMockDeploymentStore(ctrl)
	s.mockFactory.EXPECT().Deployment().AnyTimes().Return(s.mockDeplolymentStore)

	s.mockVolumeStore = store.NewMockVolumeStore(ctrl)
	s.mockFactory.EXPECT().Volume().AnyTimes().Return(s.mockVolumeStore)

	s.mockPodStore = store.NewMockPodStore(ctrl)
	s.mockFactory.EXPECT().Pod().AnyTimes().Return(s.mockPodStore)
}

func TestPod(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestPodService_GetPodByName() {
	s.mockPodStore.EXPECT().GetPodByName(gomock.Eq(context.TODO()), gomock.Eq("name_0")).Return(s.pods[0], nil)
	type fields struct {
		store    store.Factory
		redisCli redis.RedisCli
	}
	type args struct {
		ctx     context.Context
		podName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Pod
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "test case 1",
			fields: fields{store: s.mockFactory},
			args: args{
				ctx:     context.TODO(),
				podName: "name_0",
			},
			want:    s.pods[0],
			wantErr: false,
		},
	}
	for _, tt := range tests {
		podService := &PodService{store: tt.fields.store}
		pod, err := podService.GetPodByName(tt.args.ctx, tt.args.podName)
		assert.Equal(s.T(), tt.want, pod)
		assert.Equal(s.T(), tt.wantErr, err != nil)
	}
}

func (s *Suite) TestPodService_Create() {
	s.mockPodStore.EXPECT().Create(gomock.Eq(context.TODO()), gomock.Eq(s.pods[0])).Return(nil)

	type fields struct {
		store    store.Factory
		redisCli redis.RedisCli
	}
	type args struct {
		ctx context.Context
		pod *model.Pod
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "test case 1",
			fields: fields{store: s.mockFactory},
			args: args{
				ctx: context.TODO(),
				pod: s.pods[0],
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		fmt.Printf("begin to test\n")
		podService := &PodService{store: tt.fields.store}

		err := podService.Create(tt.args.ctx, tt.args.pod)
		assert.Equal(s.T(), tt.wantErr, err != nil)
	}
}

func (s *Suite) TestPodService_Update() {
	s.mockPodStore.EXPECT().Update(gomock.Eq(context.TODO()), gomock.Eq(s.pods[1])).Return(nil)

	type fields struct {
		store    store.Factory
		redisCli redis.RedisCli
	}
	type args struct {
		ctx context.Context
		pod *model.Pod
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "test case 1",
			fields: fields{store: s.mockFactory},
			args: args{
				ctx: context.TODO(),
				pod: s.pods[1],
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			podService := &PodService{
				store:    tt.fields.store,
				redisCli: tt.fields.redisCli,
			}
			if err := podService.Update(tt.args.ctx, tt.args.pod); (err != nil) != tt.wantErr {
				t.Errorf("PodService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *Suite) TestPodService_Delete() {
	s.mockPodStore.EXPECT().Delete(gomock.Eq(context.TODO()), gomock.Eq(s.pods[5]))
	type fields struct {
		store    store.Factory
		redisCli redis.RedisCli
	}
	type args struct {
		ctx context.Context
		pod *model.Pod
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			fields: fields{store: s.mockFactory},
			args: args{
				ctx: context.TODO(),
				pod: s.pods[5],
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			podService := &PodService{
				store:    tt.fields.store,
				redisCli: tt.fields.redisCli,
			}
			if err := podService.Delete(tt.args.ctx, tt.args.pod); (err != nil) != tt.wantErr {
				t.Errorf("PodService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
