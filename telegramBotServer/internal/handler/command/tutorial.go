package command

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/slider"
	"telegramCLient/internal/components"
)

type TutorialCommandHandler struct {
	Slider components.Slider
}

func NewTutorialCommandHandler() *TutorialCommandHandler {
	return &TutorialCommandHandler{
		*components.NewSlider(),
	}
}

func (c *TutorialCommandHandler) Init() []bot.Option {
	return []bot.Option{
		bot.WithMessageTextHandler("/tutorial", bot.MatchTypeExact, c.tutorialCallback),
	}
}

func (c *TutorialCommandHandler) tutorialCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	slides := []slider.Slide{
		{
			Text:  "*0\\. YouTube* is an American online video sharing and social media platform headquartered in San Bruno, California\\. It was launched on _February 14, 2005_, by *Steve Chen*, *Chad Hurley*, and *Jawed Karim*",
			Photo: "https://images.unsplash.com/photo-1733725071146-5e4157a308da?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3wxODY2Nzh8MHwxfHJhbmRvbXx8fHx8fHx8fDE3MzgyMzA1MTZ8&ixlib=rb-4.0.3&q=80&w=1080",
		},
		{
			Text:  "*1\\. VK* \\(short for its original name VKontakte; Russian: ВКонтакте, meaning InContact\\) is a Russian online social media and social networking service based in *Saint Petersburg*",
			Photo: "https://images.unsplash.com/photo-1733725071146-5e4157a308da?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3wxODY2Nzh8MHwxfHJhbmRvbXx8fHx8fHx8fDE3MzgyMzA1MTZ8&ixlib=rb-4.0.3&q=80&w=1080",
		},
		{
			Text:  "*2\\. Skype* is a proprietary telecommunications application operated by Skype Technologies, a division of *Microsoft*, best known for VoIP\\-based videotelephony, videoconferencing and voice calls",
			Photo: "https://images.unsplash.com/photo-1733725071146-5e4157a308da?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3wxODY2Nzh8MHwxfHJhbmRvbXx8fHx8fHx8fDE3MzgyMzA1MTZ8&ixlib=rb-4.0.3&q=80&w=1080",
		},
		{
			Text:  "*3\\. Reddit* \\(\\/ˈrɛdɪt\\/, stylized as reddit\\) is an American social news aggregation, web content rating, and discussion website",
			Photo: "https://images.unsplash.com/photo-1733725071146-5e4157a308da?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3wxODY2Nzh8MHwxfHJhbmRvbXx8fHx8fHx8fDE3MzgyMzA1MTZ8&ixlib=rb-4.0.3&q=80&w=1080",
		},
		{
			Text:  "*4\\. Twitter* is an American microblogging and social networking service on which users post and interact with messages known as *tweets*",
			Photo: "https://images.unsplash.com/photo-1733725071146-5e4157a308da?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3wxODY2Nzh8MHwxfHJhbmRvbXx8fHx8fHx8fDE3MzgyMzA1MTZ8&ixlib=rb-4.0.3&q=80&w=1080",
		},
		{
			Text:  "*5\\. Pinterest* is an image sharing and social media service designed to enable saving and discovery of information on the internet using images, and on a smaller scale, animated GIFs and videos, in the form of pinboards",
			Photo: "https://images.unsplash.com/photo-1733725071146-5e4157a308da?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3wxODY2Nzh8MHwxfHJhbmRvbXx8fHx8fHx8fDE3MzgyMzA1MTZ8&ixlib=rb-4.0.3&q=80&w=1080",
		},
		{
			Text:  "*6\\. Instagram* is an American photo and video sharing social networking service founded by *Kevin Systrom* and *Mike Krieger*\\. In April 2012, Facebook Inc\\. acquired the service for approximately *US$1 billion* in cash and stock",
			Photo: "https://images.unsplash.com/photo-1733725071146-5e4157a308da?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3wxODY2Nzh8MHwxfHJhbmRvbXx8fHx8fHx8fDE3MzgyMzA1MTZ8&ixlib=rb-4.0.3&q=80&w=1080",
		},
		{
			Text:  "*7\\. LinkedIn* is an American business and employment\\-oriented online service that operates via websites and mobile apps\\. Launched on May 5, 2003",
			Photo: "https://images.unsplash.com/photo-1733725071146-5e4157a308da?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3wxODY2Nzh8MHwxfHJhbmRvbXx8fHx8fHx8fDE3MzgyMzA1MTZ8&ixlib=rb-4.0.3&q=80&w=1080",
		},
		{
			Text:  "*8\\. Facebook* is an American online social media and social networking service owned by Meta Platforms\\. Founded in 2004 by *Mark Zuckerberg* with fellow Harvard College students and roommates *Eduardo Saverin*, *Andrew McCollum*, *Dustin Moskovitz*, and *Chris Hughes*",
			Photo: "https://images.unsplash.com/photo-1733725071146-5e4157a308da?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3wxODY2Nzh8MHwxfHJhbmRvbXx8fHx8fHx8fDE3MzgyMzA1MTZ8&ixlib=rb-4.0.3&q=80&w=1080",
		},
	}
	c.Slider.CreateAndRun(ctx, b, update, slides)
}
