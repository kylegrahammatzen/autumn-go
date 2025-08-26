# Examples

---

This directory contains working examples demonstrating how to use the Autumn Go SDK for subscription and usage tracking.

## Available Examples

---

### [Basic Example](./basic/)

Complete demonstration of the core Autumn SDK functionality including:
- Customer attachment to products
- Feature access checking with balance information  
- Usage tracking and consumption
- Automated checkout flow for upgrades

## Getting Started

---

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Set your API key in the example files (replace the test key)

3. Configure your products in the Autumn dashboard to match the example IDs

4. Run an example:
   ```bash
   cd basic
   go run main.go
   ```

## Common Patterns

---

All examples demonstrate the standard Autumn workflow:

1. Attach - Connect customers to subscription plans
2. Check - Verify access before granting features  
3. Track - Record usage when features are consumed
4. Checkout - Handle upgrades when limits are reached

## Need Help?

---

See the [main documentation](../README.md) for installation and setup instructions.