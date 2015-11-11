package helpers

import (
  "snaga-team/models"
)

func SanitizeNewEvent(event *models.Event) {
  event.CancellationDate = ""
  event.LastUpdateDate = ""

}

func SanitizeEvent(newData *models.Event, storedData *models.Event) {
  event.CancellationDate = storedData.CancellationDate
  event.LastUpdateDate = storedData.LastUpdateDate
  event.CreatorId = storedData.CreatorId
}
