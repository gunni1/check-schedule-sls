# How To
## Environment Variables

| Variable | Description |
| --- | --- |
| TEACHER_CODE | Acronym for teacher in schedule API |
| PAGE_USER | Username for schedule web API |
| PAGE_PASSWORD | Password for schedule web API |
| NOTIFICATION_TARGET | Chat ID of the telegram conversation with the notification bot |
| DAYS_COUNT | Number of future days to check the schedule. Starting from today, e.g. 2 means tomorrow and the day after. Weedend is skipped.
| BOT_TOKEN | Bot-Token of the notification bot called `StundenplanInfoBot` | 

## local testing
`sls invoke local -f check-schedule -e TEACHER_CODE=BD -e PAGE_PW=<SCHEDULE_API_PW> -e PAGE_USER=<SCHEDULE_API_USER> -e NOTIFICATION_TARGET=<BOT_CHAT_ID> -e BOT_TOKEN=<TOKEN>`

# TODO
- Aus einer gefundenen änderung einen hash berechnen und in eine db speichern
- vor benachrichtigung über änderung hash in db prüfen und wenn bereits vorhanden, nicht benachrichtigen
- bereits übermittelte benachrichtigungen aufräumen (zeitstempel)