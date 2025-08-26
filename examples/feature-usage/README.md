# Feature Usage Management Example

Complete example demonstrating feature usage tracking and entity management using the Autumn Go SDK.

## Setup

1. Get your API key from https://app.useautumn.com/sandbox/dev
2. Replace the example key in `main.go` with your actual key
3. Run the example:
   ```bash
   cd examples/feature-usage
   go run main.go
   ```

## Features Demonstrated

- **Direct Usage Setting**: Set feature usage values directly via API
- **Entity Management**: Create and manage entities like seats, team members
- **Batch Operations**: Create multiple entities in a single request
- **Usage Tracking**: Track different types of usage (API calls, storage, etc.)
- **Entity Cleanup**: Remove entities when no longer needed

## Output

```
Cleaning up existing entities...
Cleaned up entity: seat_001
Cleaned up entity: seat_002
Cleaned up entity: seat_003

Setting feature usage...
Set usage for 'api_calls' to 150 for customer jane_doe

Creating entities...
Created entity: seat_001 (John Doe)
Created 2 additional entities

Setting storage usage...
Set storage usage to 25.5 GB

Removing team member...
Deleted entity: seat_003 (Bob Johnson)

Feature usage management completed successfully!

Summary:
- Set API calls usage to 150
- Created 3 team member seats
- Set storage usage to 25.5 GB
- Removed 1 team member (2 seats remaining)
```

## API Endpoints Used

- `POST /usage` - Set feature usage values
- `POST /customers/{customer_id}/entities` - Create entities
- `DELETE /customers/{customer_id}/entities/{entity_id}` - Delete entities

## Use Cases

- **SaaS Seat Management**: Track team members, assign seats
- **Usage Quotas**: Monitor API calls, storage, bandwidth usage  
- **Resource Tracking**: Track any measurable feature consumption
- **Team Changes**: Handle onboarding/offboarding of users