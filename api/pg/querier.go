// Code generated by sqlc. DO NOT EDIT.

package pg

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateLabelColor(ctx context.Context, arg CreateLabelColorParams) (LabelColor, error)
	CreateOrganization(ctx context.Context, arg CreateOrganizationParams) (Organization, error)
	CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error)
	CreateProjectLabel(ctx context.Context, arg CreateProjectLabelParams) (ProjectLabel, error)
	CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) (RefreshToken, error)
	CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error)
	CreateTaskAssigned(ctx context.Context, arg CreateTaskAssignedParams) (TaskAssigned, error)
	CreateTaskGroup(ctx context.Context, arg CreateTaskGroupParams) (TaskGroup, error)
	CreateTaskLabelForTask(ctx context.Context, arg CreateTaskLabelForTaskParams) (TaskLabel, error)
	CreateTeam(ctx context.Context, arg CreateTeamParams) (Team, error)
	CreateUserAccount(ctx context.Context, arg CreateUserAccountParams) (UserAccount, error)
	DeleteExpiredTokens(ctx context.Context) error
	DeleteRefreshTokenByID(ctx context.Context, tokenID uuid.UUID) error
	DeleteRefreshTokenByUserID(ctx context.Context, userID uuid.UUID) error
	DeleteTaskAssignedByID(ctx context.Context, arg DeleteTaskAssignedByIDParams) (TaskAssigned, error)
	DeleteTaskByID(ctx context.Context, taskID uuid.UUID) error
	DeleteTaskGroupByID(ctx context.Context, taskGroupID uuid.UUID) (int64, error)
	DeleteTasksByTaskGroupID(ctx context.Context, taskGroupID uuid.UUID) (int64, error)
	DeleteTeamByID(ctx context.Context, teamID uuid.UUID) error
	GetAllOrganizations(ctx context.Context) ([]Organization, error)
	GetAllProjects(ctx context.Context) ([]Project, error)
	GetAllProjectsForTeam(ctx context.Context, teamID uuid.UUID) ([]Project, error)
	GetAllTaskGroups(ctx context.Context) ([]TaskGroup, error)
	GetAllTasks(ctx context.Context) ([]Task, error)
	GetAllTeams(ctx context.Context) ([]Team, error)
	GetAllUserAccounts(ctx context.Context) ([]UserAccount, error)
	GetAssignedMembersForTask(ctx context.Context, taskID uuid.UUID) ([]TaskAssigned, error)
	GetLabelColorByID(ctx context.Context, labelColorID uuid.UUID) (LabelColor, error)
	GetLabelColors(ctx context.Context) ([]LabelColor, error)
	GetProjectByID(ctx context.Context, projectID uuid.UUID) (Project, error)
	GetProjectLabelByID(ctx context.Context, projectLabelID uuid.UUID) (ProjectLabel, error)
	GetProjectLabelsForProject(ctx context.Context, projectID uuid.UUID) ([]ProjectLabel, error)
	GetRefreshTokenByID(ctx context.Context, tokenID uuid.UUID) (RefreshToken, error)
	GetTaskByID(ctx context.Context, taskID uuid.UUID) (Task, error)
	GetTaskGroupByID(ctx context.Context, taskGroupID uuid.UUID) (TaskGroup, error)
	GetTaskGroupsForProject(ctx context.Context, projectID uuid.UUID) ([]TaskGroup, error)
	GetTaskLabelsForTaskID(ctx context.Context, taskID uuid.UUID) ([]TaskLabel, error)
	GetTasksForTaskGroupID(ctx context.Context, taskGroupID uuid.UUID) ([]Task, error)
	GetTeamByID(ctx context.Context, teamID uuid.UUID) (Team, error)
	GetTeamsForOrganization(ctx context.Context, organizationID uuid.UUID) ([]Team, error)
	GetUserAccountByID(ctx context.Context, userID uuid.UUID) (UserAccount, error)
	GetUserAccountByUsername(ctx context.Context, username string) (UserAccount, error)
	UpdateTaskDescription(ctx context.Context, arg UpdateTaskDescriptionParams) (Task, error)
	UpdateTaskGroupLocation(ctx context.Context, arg UpdateTaskGroupLocationParams) (TaskGroup, error)
	UpdateTaskLocation(ctx context.Context, arg UpdateTaskLocationParams) (Task, error)
	UpdateTaskName(ctx context.Context, arg UpdateTaskNameParams) (Task, error)
}

var _ Querier = (*Queries)(nil)
