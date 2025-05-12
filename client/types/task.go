package types

type Task struct {
    TaskId        string        `json:"taskId"`
    ResourceId    string        `json:"resourceId"`
}

type TaskInfo struct {
    Id            string           `json:"id"`
    Name          string           `json:"name"`
    Detail        string           `json:"detail"`
    State         string           `json:"state"`
    StartTime     string           `json:"startTime"`
    EndTime       string           `json:"endTime"`
    ActorName     string           `json:"actorName"`
    Error         string           `json:"error"`
    ErrorCode     string           `json:"errorCode"`
    Cancelable    bool             `json:"cancelable"`
    Canceled      bool             `json:"canceled"`
    TargetName    string           `json:"targetName"`
    TargetId      string           `json:"targetId"`
    TargetType    string           `json:"targetType"`
    Events        []interface{}    `json:"events"`
    ProcessId     string           `json:"processId"`
    Progress      int              `json:"progress"`
    ChildTasks    []TaskInfo       `json:"childTasks"`
}