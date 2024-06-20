package notifications

import (
	"context"
	"time"
)

func (n *Notification) Run(ctx context.Context) error {
	checkBirth := time.After(0)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-checkBirth:
			notifications, err := n.srvEmployee.TodayBirthdays(ctx)
			if err != nil {
				return err
			}

			for _, notification := range notifications {
				if err = n.srvNotification.Notice(ctx,
					notification.Subscriber.FullName,
					notification.Subscriber.Email,
					notification.Celebrator.FullName,
					notification.Celebrator.Email); err != nil {
					n.logger.Error("failed to notice", "subscriber", notification.Subscriber.FullName, "celebrator", notification.Celebrator.FullName)
				}
			}

			checkBirth = n.timeAfter()
		}
	}
}
