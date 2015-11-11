package helpers

import (
  "snaga-team/models"
)

func SanitizeNewEvent(event *models.Event) {
  event.CancellationDate = ""
  event.LastUpdateDate = ""

}

func SanitizeEvent(newData *models.Event, storedData *models.Event) {
  newData.CancellationDate = storedData.CancellationDate
  newData.LastUpdateDate = storedData.LastUpdateDate
  newData.CreatorId = storedData.CreatorId
}
