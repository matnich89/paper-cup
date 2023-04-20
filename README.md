# Video API 

## Assumptions 

- Videos must be uploaded / stored before any annotations 
  are created for it... so the video model does not have an annotations slice 
- Related to not having any annotations... there seemed to be no requirement for Get Videos end point
  felt this was slightly unusual. but as it's not required decided to save myself some time :) 
- There was no specific requirements around uniqueness of records ao I decided to use my own judgement 
  on what unique should look like for the records (see the migration SQL )
- If a video is deleted then we want to cascade deletion of any annotations related to it
- Once an annotation has been created it can't be switched to relate to another video
- Can only have one video for a URL in our database
- We are only working with hours:minutes:seconds for videos

## Things missing / I am not proud of

### Given time constraints there are various things that could be improved

- The User password storage... I would never store a password in plain text usually, 
  but given users were not a hard requirement I put something together quickly 
  to enable the JWT authentication flow

- Very little Handler tests only about 35% coverage (these were proving very time-consuming to create)  I need to
 look at simplifying the approach I took to these. In a prod env I wouldn't allow this in with such little handler coverage.
- No middleware tests
  
- No validation of payloads... I'd generally use `https://github.com/go-playground/validator`
  for validation in production
- Lack of logging... in a production env there would be far more logging...sorry
- No linter setup (though go linter has been ran)

## Notes

- The application can simply be started from the root with `docker-compose up`

- Please note the urls provided for videos must start with the http:// protocol defination 
  to pass validation

- I have included a very basic Postman collection for the apis (see `postman-examples.json`

- I would usually use `go mod vendor` to vendor the dependencies, but decided to leave it out 
  for this

- Its worth mentioning that the database tests would not run on windows
it seems there is some weird issue with invoking Docker Desktop...so Macs and Linux only please :) 

