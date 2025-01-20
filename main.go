package main

import (
	"github.com/gocql/gocql"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/datasource/scylladb"
)

type Vote struct {
	PollID gocql.UUID `json:"poll_id"`
	VoteID gocql.UUID `json:"vote_id"`
	Option string     `json:"option"`
}

type Poll struct {
	ID       gocql.UUID `json:"id"`
	Question string     `json:"question"`
	Options  []string   `json:"options"`
}

type PollService struct {
	ScyllaDB *scylladb.Client
}

func (s *PollService) CreatePoll(ctx *gofr.Context, poll *Poll) error {

	poll.ID = gocql.TimeUUID()

	query := `INSERT INTO polls (poll_id, question, options) VALUES (?, ?, ?)`
	err := s.ScyllaDB.ExecWithCtx(ctx, query, poll.ID, poll.Question, poll.Options)
	if err != nil {
		return err
	}
	return nil
}

func (s *PollService) DeletePoll(ctx *gofr.Context, pollID gocql.UUID) error {
	query := `DELETE FROM polls WHERE poll_id = ?`
	err := s.ScyllaDB.ExecWithCtx(ctx, query, pollID)
	if err != nil {
		return err
	}
	return nil
}

type VoteService struct {
	ScyllaDB *scylladb.Client
}

func (s *VoteService) VoteOnPoll(ctx *gofr.Context, vote *Vote) error {
	vote.VoteID = gocql.TimeUUID()

	query := `INSERT INTO votes (poll_id, vote_id, option) VALUES (?, ?, ?)`
	err := s.ScyllaDB.ExecWithCtx(ctx, query, vote.PollID, vote.VoteID, vote.Option)
	if err != nil {
		return err
	}
	return nil
}

func (s *PollService) DeleteVotesForPoll(ctx *gofr.Context, pollID gocql.UUID) error {
	query := `DELETE FROM votes WHERE poll_id = ?`
	err := s.ScyllaDB.ExecWithCtx(ctx, query, pollID)
	if err != nil {
		return err
	}
	return nil
}

func (s *PollService) UpdatePoll(ctx *gofr.Context, pollID gocql.UUID, updatedPoll *Poll) error {
	query := `UPDATE polls SET question = ?, options = ? WHERE poll_id = ?`
	err := s.ScyllaDB.ExecWithCtx(ctx, query, updatedPoll.Question, updatedPoll.Options, pollID)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	app := gofr.New()
	client := scylladb.New(scylladb.Config{
		Host:     "localhost",
		Keyspace: "polls_db",
		Port:     2025,
		Username: "root",
		Password: "password",
	})
	app.AddScyllaDB(client)
	pollService := &PollService{ScyllaDB: client}

	app.POST("/polls", func(c *gofr.Context) (interface{}, error) {
		var newPoll Poll
		err := c.Bind(&newPoll)
		if err != nil {
			return nil, err
		}

		err = pollService.CreatePoll(c, &newPoll)
		if err != nil {
			return nil, err
		}

		return newPoll, nil
	})
	voteService := &VoteService{ScyllaDB: client}
	app.POST("/polls/{poll_id}/vote", func(c *gofr.Context) (interface{}, error) {
		pollID := c.PathParam("poll_id")
		parsedPollID, err := gocql.ParseUUID(pollID)
		if err != nil {
			return nil, err
		}

		var newVote Vote
		newVote.PollID = parsedPollID
		err = c.Bind(&newVote)
		if err != nil {
			return nil, err
		}

		err = voteService.VoteOnPoll(c, &newVote)
		if err != nil {
			return nil, err
		}

		return newVote, nil
	})
	app.DELETE("/polls/{poll_id}", func(c *gofr.Context) (interface{}, error) {
		pollID := c.PathParam("poll_id")

		parsedID, err := gocql.ParseUUID(pollID)
		if err != nil {
			return nil, err
		}

		err = pollService.DeletePoll(c, parsedID)
		if err != nil {
			return nil, err
		}

		err = pollService.DeleteVotesForPoll(c, parsedID)
		if err != nil {
			return nil, err
		}

		return "Poll and associated votes deleted successfully", nil
	})

	app.PUT("/polls/{poll_id}", func(c *gofr.Context) (interface{}, error) {
		pollID := c.PathParam("poll_id")
		parsedPollID, err := gocql.ParseUUID(pollID)
		if err != nil {
			return nil, err
		}
		var updatedPoll Poll
		err = c.Bind(&updatedPoll)
		if err != nil {
			return nil, err
		}

		err = pollService.UpdatePoll(c, parsedPollID, &updatedPoll)
		if err != nil {
			return nil, err
		}

		return map[string]string{
			"message": "Poll updated successfully",
		}, nil
	})

	app.Run()
}
