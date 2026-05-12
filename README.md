# Smart Parking Microservices

Smart Parking is a simple microservice-based parking management prototype. The project is designed for deployment to Google Kubernetes Engine and satisfies the requirement to create at least three stub microservices that write data into one shared database.

The system models a parking application where users can view parking spots, reserve a spot for a time window, check in and check out, and receive a bill for the parking session.

## Core Idea

The application consists of three services:

1. Spot Service
2. Reservation Service
3. Billing Service

Each service has its own API and responsibility area, but all services persist data into the same database.

## Architecture

```text
Client / API testing tool
        |
        | HTTP
        v
+-------------------+     +-------------------------+
| Spot Service      | --> |                         |
+-------------------+     |                         |
                          |                         |
+-------------------+     |                         |
| Reservation       | --> |  Shared PostgreSQL DB   | 
| Service           |     |                         |
+-------------------+     |                         |
                          |                         |
+-------------------+     |                         |
| Billing Service   | --> |                         |
+-------------------+     +-------------------------+
```

## Services

### 1. Spot Service

The Spot Service manages parking infrastructure.

Responsibilities:

- Manage parking zones.
- Manage parking spots.
- Store spot type.
- Store and update spot availability.
- Provide available spots for reservation.

Main domain concepts:

- Parking zone: logical parking area, for example `Zone A`, `Zone B`, `Underground`.
- Parking spot: concrete parking place, for example `A-01`.
- Spot type: `REGULAR`, `EV`, `DISABLED`.
- Spot status: `AVAILABLE`, `RESERVED`, `OCCUPIED`, `MAINTENANCE`.

Possible API endpoints:

```http
POST  /spots
GET   /spots
GET   /spots/available
GET   /spots/{spotId}
PATCH /spots/{spotId}/status
```

Example spot:

```json
{
  "id": "spot-1",
  "zone": "Zone A",
  "number": "A-01",
  "type": "EV",
  "status": "AVAILABLE"
}
```

Domain rules:

- A spot can be reserved only if its status is `AVAILABLE`.
- A reserved spot should not be offered as available.
- A spot with status `MAINTENANCE` cannot be reserved.
- During check-in, a reserved spot can become `OCCUPIED`.
- During check-out, an occupied spot can become `AVAILABLE` again.

### 2. Reservation Service

The Reservation Service manages user reservations and parking sessions.

Responsibilities:

- Create a reservation for a specific user and spot.
- Store reservation start and end time.
- Track reservation status.
- Handle check-in.
- Handle check-out.
- Store reservation data in the shared database.

Main domain concepts:

- Reservation: user's booking of a parking spot for a time window.
- Check-in: user starts using the reserved spot.
- Check-out: user finishes using the reserved spot.

Reservation statuses:

- `CREATED`
- `CHECKED_IN`
- `CHECKED_OUT`
- `CANCELLED`

Possible API endpoints:

```http
POST  /reservations
GET   /reservations
GET   /reservations/{reservationId}
GET   /reservations?userId={userId}
PATCH /reservations/{reservationId}/check-in
PATCH /reservations/{reservationId}/check-out
PATCH /reservations/{reservationId}/cancel
```

Example reservation:

```json
{
  "id": "reservation-1",
  "userId": "user-1",
  "spotId": "spot-1",
  "startTime": "2026-05-12T10:00:00Z",
  "endTime": "2026-05-12T12:00:00Z",
  "status": "CREATED"
}
```

Domain rules:

- A reservation should reference an existing parking spot.
- A reservation should be created only for an available spot.
- After reservation creation, the related spot can be marked as `RESERVED`.
- Only a `CREATED` reservation can be checked in.
- Only a `CHECKED_IN` reservation can be checked out.
- A checked-out reservation can be used to create a billing record.
- A cancelled reservation should not be billed.

### 3. Billing Service

The Billing Service manages bills and payment status stubs.

Responsibilities:

- Create a bill for a reservation.
- Store user ID and reservation ID on the bill.
- Calculate payment amount using parking duration.
- Return bills by user ID.
- Return bills by reservation ID.
- Mark a specific bill as paid.
- Store billing data in the shared database.

Main domain concepts:

- Bill: payment record for a reservation.
- Payment status: current payment state of a bill.

Bill statuses:

- `PENDING`
- `PAID`
- `FAILED`
- `CANCELLED`

Possible API endpoints:

```http
POST  /billing
GET   /billing
GET   /billing/{billId}
GET   /billing?userId={userId}
GET   /billing?reservationId={reservationId}
GET   /billing?status={status}
PATCH /billing/{billId}/pay
```

Domain rules:

- A bill should reference a reservation.
- A reservation should normally have no more than one active bill.
- A newly created bill has status `PENDING`.
- Only a `PENDING` bill can be marked as `PAID`.
- Payment is a stub operation; no real payment provider is used.
- The amount can be calculated with a simple hourly rate.

## Shared Database Model

The database can contain separate tables for each service area.

Suggested tables:

- `parking_zones`
- `spots`
- `reservations`
- `bills`

Suggested ownership:

- Spot Service writes to `parking_zones` and `spots`.
- Reservation Service writes to `reservations`.
- Billing Service writes to `bills`.

## Example User Flow

1. Spot Service creates parking zones and spots.
2. User requests available spots.
3. Reservation Service creates a reservation for an available spot.
4. Spot Service marks the spot as `RESERVED`.
5. User checks in.
6. Reservation Service marks the reservation as `CHECKED_IN`.
7. Spot Service marks the spot as `OCCUPIED`.
8. User checks out.
9. Reservation Service marks the reservation as `CHECKED_OUT`.
10. Spot Service marks the spot as `AVAILABLE`.
11. Billing Service creates a bill for the reservation.
12. User pays the bill.
13. Billing Service marks the bill as `PAID`.