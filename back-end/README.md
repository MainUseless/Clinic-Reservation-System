| endpoint                        | input                                                          | return                                  |
| ------------------------------- | -------------------------------------------------------------- | --------------------------------------- |
| POST /api/account/signin       | Query params : email, password                                 | error or JWT token as json              |
| POST /api/account/signup        | Query params : email, password, type( doctor or patient), name | error or message as json                |
| POST /api/doctor/appointment    | Barer token : JWT & Query params : timestamp                   | JSON result ture or JSON error          |
| GET /api/doctor/appointment     | Barer token : JWT                                              | unauthorized or appointments array json |
| DELETE /api/doctor/appointment  | TODO                                                           | TODO                                    |
| POST /api/patient/appointment   | Barer token : JWT & Query params : new_timestamp              | JSON result ture or JSON error          |
| PUT /api/patient/appointment    | Barer token : JWT & Query params : old_id, new_timestamp       | JSON result ture or JSON error          |
| GET /api/patient/appointment    | Barer token : JWT                                              | unauthorized or appointments array json |
| DELETE /api/patient/appointment | Barer token : JWT & Query params : appointment_id              | JSON result ture or JSON error          |

note : JWT never expires :)

timestamp format : 2001-02-02.14:24
