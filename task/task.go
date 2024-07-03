package task

import (
	"context"
	"github.com/ics-sigs/ics-go-sdk/client/methods"
	"github.com/ics-sigs/ics-go-sdk/client/types"
	"sync"
	"time"
)

func (t *TaskService) GetTaskInfo(ctx context.Context, task *types.Task) (*types.TaskInfo, error) {
	taskInfo, err := methods.GetTaskInfo(ctx, t.RestAPITripper, task)
	return taskInfo, err
}

func (t *TaskService) WaitForResult(ctx context.Context, task *types.Task) (taskInfo *types.TaskInfo, err error) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		count := 1
		for {
			taskInfo, err = methods.GetTaskInfo(ctx, t.RestAPITripper, task)
			if err != nil || count > 72000 {
				break
			}
			if taskInfo == nil || taskInfo.State == "FINISHED" || taskInfo.State == "ERROR" {
				break
			}
			time.Sleep(100 * time.Millisecond)
			count++
		}
	}()
	wg.Wait()
	return taskInfo, err
}
