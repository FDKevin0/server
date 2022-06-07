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

package handler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/bangumi/server/internal/domain"
	"github.com/bangumi/server/internal/errgo"
	"github.com/bangumi/server/internal/model"
	"github.com/bangumi/server/internal/web/res"
	"github.com/bangumi/server/internal/web/util"
)

func (h Handler) getTopic(c *fiber.Ctx, topicType domain.TopicType, id model.TopicIDType) (model.Topic, error) {
	page, err := getPageQuery(c, defaultPageLimit, defaultMaxPageLimit)
	if err != nil {
		return model.Topic{}, c.Status(http.StatusNotFound).JSON(res.Error{
			Title:   "Not Found",
			Details: util.DetailFromRequest(c),
		})
	}

	topic, err := h.t.Get(c.Context(), topicType, id, page.Limit, page.Offset)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return model.Topic{}, c.Status(http.StatusNotFound).JSON(res.Error{
				Title:   "Not Found",
				Details: util.DetailFromRequest(c),
			})
		}
		return model.Topic{}, errgo.Wrap(err, "Topic.Get")
	}
	return topic, nil
}

func (h Handler) getUserMapOfTopics(c *fiber.Ctx, topics ...model.Topic) (map[uint32]model.User, error) {
	userIDs := make([]model.UIDType, 0)
	for _, topic := range topics {
		userIDs = append(userIDs, topic.UID)
		for _, v := range topic.Comments.Data {
			userIDs = append(userIDs, v.UID)
		}
	}
	userMap, err := h.u.GetByIDs(c.Context(), dedupeUIDs(userIDs...)...)
	if err != nil {
		return nil, errgo.Wrap(err, "user.GetByIDs")
	}
	return userMap, nil
}

func (h Handler) listTopics(c *fiber.Ctx, topicType domain.TopicType, id uint32) error {
	page, err := getPageQuery(c, defaultPageLimit, defaultMaxPageLimit)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(res.Error{
			Title:   "Not Found",
			Details: util.DetailFromRequest(c),
		})
	}
	var response = res.Paged{
		Limit:  page.Limit,
		Offset: page.Offset,
	}

	topics, err := h.t.ListTopics(c.Context(), topicType, id, page.Limit, page.Offset)
	if err != nil {
		return errgo.Wrap(err, "repo.topic.GetTopics")
	}

	userMap, err := h.getUserMapOfTopics(c, topics...)
	if err != nil {
		return errgo.Wrap(err, "user.GetByIDs")
	}
	count, err := h.t.Count(c.Context(), topicType, id)
	if err != nil {
		return errgo.Wrap(err, "repo.topic.Count")
	}
	response.Total = count
	var data = make([]res.Topic, len(topics))
	for i, topic := range topics {
		creator := userMap[topic.UID]
		data[i] = res.Topic{
			ID:        topic.ID,
			Title:     topic.Title,
			CreatedAt: topic.CreatedAt,
			Creator: res.Creator{
				Username: creator.UserName,
				Nickname: creator.NickName,
			},
			Replies: topic.Replies,
		}
	}
	response.Data = data
	return c.JSON(response)
}

func (h Handler) getResTopic(c *fiber.Ctx, topic model.Topic) error {
	userMap, err := h.getUserMapOfTopics(c, topic)
	if err != nil {
		return err
	}

	creator := userMap[topic.UID]
	topic.Comments.Data = model.ConvertModelCommentsToTree(topic.Comments.Data, 0)
	response := res.Topic{
		ID:        topic.ID,
		Title:     topic.Title,
		CreatedAt: topic.CreatedAt,
		Creator: res.Creator{
			Username: creator.UserName,
			Nickname: creator.NickName,
		},
		Replies: topic.Replies,
		Comments: &res.Comments{
			HasMore: topic.Comments.HasMore,
			Limit:   topic.Comments.Limit,
			Offset:  topic.Comments.Offset,
			Data:    convertModelTopicComments(topic.Comments.Data, userMap),
		},
	}
	return c.JSON(response)
}

func convertModelTopicComments(comments []model.Comment, userMap map[uint32]model.User) []res.Comment {
	replies := make([]res.Comment, 0)
	for _, v := range comments {
		creator := userMap[v.UID]
		replies = append(replies, res.Comment{
			ID:        v.ID,
			Text:      v.Content,
			CreatedAt: v.CreatedAt,
			Creator: res.Creator{
				Username: creator.UserName,
				Nickname: creator.NickName,
			},
			Replies: convertModelTopicComments(v.Replies, userMap),
		})
	}
	return replies
}
