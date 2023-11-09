| endpoint                        | input                                                        | return                                           |
| ------------------------------- | ------------------------------------------------------------ | ------------------------------------------------ |
| POST /api/account/signin       | JSON body : email, password                                 | error or JWT token as json                       |
| POST /api/account/signup        | JSON body : email, password, type( doctor or patient), name  | error or message as json                         |
| POST /api/doctor/appointment    | Barer token : JWT & Query params : timestamp                 | unauthorized or JSON result ture or JSON error  |
| GET /api/doctor/appointment     | Barer token : JWT                                            | unauthorized or appointments array json          |
| DELETE /api/doctor/appointment  | Barer token : JWT & Query params : appointment_id            | unauthorized or JSON result                      |
| POST /api/patient/appointment   | Barer token : JWT & Query params : timestamp                | unauthorized or JSON result ture or JSON error   |
| PUT /api/patient/appointment    | Barer token : JWT & Query params : appointment_id, timestamp | unauthorized or JSON result ture or JSON error   |
| GET /api/patient/appointment    | Barer token : JWT                                            | unauthorized or patients appointments array json |
| GET /api/patient/appointment    | Barer token : JWT & doctor_id                                | unauthorized or doctors appointments array json  |
| DELETE /api/patient/appointment | Barer token : JWT & Query params : appointment_id            | unauthorized or JSON result ture or JSON error   |
| GET /api/patient/doctors        | Barer token : JWT                                            | unauthorized or JSON with doctors names and ids  |

note : JWT never expires :)

timestamp format : 2001-02-02 14:24
