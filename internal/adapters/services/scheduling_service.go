package services

import (
	"context"
	"errors"

	cloudScheduler "cloud.google.com/go/scheduler/apiv1"
	schedulerpb "cloud.google.com/go/scheduler/apiv1/schedulerpb"
)

type ScheduleService struct {
	client *cloudScheduler.CloudSchedulerClient
}

func NewScheduleService() *ScheduleService {
	client, err := cloudScheduler.NewCloudSchedulerClient(context.Background())
	if err != nil {
		panic(err)
	}

	return &ScheduleService{
		client: client,
	}
}

func (ss *ScheduleService) ScheduleJob(uri string, name string) error {
	// Define the job
	req := &schedulerpb.CreateJobRequest{
		Parent: "projects/<your-project-id>/locations/us-central1",
		Job: &schedulerpb.Job{
			Target: &schedulerpb.Job_HttpTarget{
				HttpTarget: &schedulerpb.HttpTarget{
					Uri:        uri,
					HttpMethod: schedulerpb.HttpMethod_POST,
				},
			},
			Schedule: "0 0 * * *", // Cron syntax for daily scheduling
			Name: name,
		},
	}

	_, err := ss.client.CreateJob(context.Background(), req)
	if err != nil {
		return errors.New("failed to create job")
	}

	return nil
}