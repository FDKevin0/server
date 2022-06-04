// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package comment_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/bangumi/server/internal/comment"
	"github.com/bangumi/server/internal/dal/dao"
	"github.com/bangumi/server/internal/dal/query"
	"github.com/bangumi/server/internal/domain"
	"github.com/bangumi/server/internal/model"
	"github.com/bangumi/server/internal/test"
)

func getRepo(t *testing.T) domain.CommentRepo {
	t.Helper()
	repo, err := comment.NewMysqlRepo(query.Use(test.GetGorm(t)), zap.NewNop())
	require.NoError(t, err)

	return repo
}

func TestGet(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)

	s, err := repo.Get(context.Background(), model.CommentIndex, 1038)
	require.NoError(t, err)

	require.Equal(t, model.CommentIDType(1038), s.ID)
}

func TestMysqlRepo_GetCommentsByMentionedID(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)

	s, err := repo.GetCommentsByMentionedID(context.Background(), model.CommentTypeSubjectTopic, 10, 10, 2)
	require.NoError(t, err)

	require.Equal(t, model.CommentIDType(1), s.Data[1].ID)

	require.Equal(t, model.CommentIDType(2), s.Data[2].ID)
}

func TestMysqlRepo_ConvertDao(t *testing.T) {
	t.Parallel()

	s, err := comment.ConvertDao(&dao.SubjectTopicComment{
		ID: 10,
	})
	require.NoError(t, err)
	require.Equal(t, s.ID, model.CommentIDType(10))
}