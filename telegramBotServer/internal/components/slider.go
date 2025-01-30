package components

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/slider"
	"strconv"
)

type Slider struct {
}

func NewSlider() *Slider {
	return &Slider{}
}

func (sl *Slider) CreateAndRun(ctx context.Context, b *bot.Bot, update *models.Update, slides []slider.Slide) {
	opts := []slider.Option{
		slider.OnSelect("Select", true, sl.sliderOnSelect),
		slider.OnCancel("Cancel", true, sl.sliderOnCancel),
	}
	newSlider := slider.New(b, slides, opts...)
	newSlider.Show(ctx, b, strconv.Itoa(int(update.Message.Chat.ID)))
}

func (sl *Slider) sliderOnSelect(ctx context.Context, b *bot.Bot, message models.MaybeInaccessibleMessage, item int) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: message.Message.Chat.ID,
		Text:   "Select " + strconv.Itoa(item),
	})
}

func (sl *Slider) sliderOnCancel(ctx context.Context, b *bot.Bot, message models.MaybeInaccessibleMessage) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: message.Message.Chat.ID,
		Text:   "Cancel",
	})
}
