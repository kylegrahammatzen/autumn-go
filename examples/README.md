# Examples

This directory contains working examples demonstrating how to use the Autumn Go SDK for subscription and usage tracking.

### [Basic Example](./basic/)

Complete demonstration of the core Autumn SDK functionality including:
- Customer attachment to products
- Feature access checking with balance information  
- Usage tracking and consumption
- Automated checkout flow for upgrades

### [Customer Management](./customer-management/)

Customer lifecycle management operations including:
- Creating and updating customer profiles
- Retrieving detailed customer information
- Setting feature balances manually
- Generating billing portal URLs for self-service

### [Product Management](./product-management/)

Product creation and configuration including:
- Creating products with pricing and feature items
- Retrieving detailed product information
- Configuring free trials and usage models
- Managing different feature types (single-use, continuous-use)

### [Feature Usage Management](./feature-usage/)

Feature usage tracking and entity management including:
- Setting feature usage values directly
- Creating and managing entities (seats, team members)
- Batch entity operations
- Usage tracking for different feature types

### [Configuration-Based Management](./config/)

Define products and features using Go configuration files including:
- Type-safe feature and product definitions
- Configuration-driven product creation
- Seamless conversion to Autumn API calls
- Similar to TypeScript CLI approach