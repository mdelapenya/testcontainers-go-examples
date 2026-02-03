// Playwright API tests for QuickPizza
const { test, expect } = require('@playwright/test');

const BASE_URL = process.env.FRONTEND_URL || 'http://localhost:3333';

test.describe('QuickPizza API Tests', () => {
  test('should return 200 on homepage', async ({ request }) => {
    const response = await request.get(BASE_URL);
    expect(response.status()).toBe(200);
  });

  test('should successfully order a pizza with authentication', async ({ request }) => {
    const restrictions = {
      maxCaloriesPerSlice: 500,
      mustBeVegetarian: false,
      maxNumberOfToppings: 6,
      minNumberOfToppings: 2
    };

    const response = await request.post(`${BASE_URL}/api/pizza`, {
      data: restrictions,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'token abcdef0123456789',
      },
    });

    expect(response.ok()).toBeTruthy();
    expect(response.status()).toBe(200);
    
    const pizza = await response.json();
    expect(pizza).toHaveProperty('pizza');
    expect(pizza.pizza).toHaveProperty('name');
    expect(pizza.pizza).toHaveProperty('ingredients');
    
    console.log(`✅ Ordered: ${pizza.pizza.name} (${pizza.pizza.ingredients.length} ingredients)`);
  });

  test('should handle multiple concurrent pizza orders', async ({ request }) => {
    const restrictions = {
      maxCaloriesPerSlice: 500,
      mustBeVegetarian: false,
      maxNumberOfToppings: 6,
      minNumberOfToppings: 2
    };

    // Make 10 concurrent requests to simulate load
    const promises = [];
    for (let i = 0; i < 10; i++) {
      promises.push(
        request.post(`${BASE_URL}/api/pizza`, {
          data: restrictions,
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'token abcdef0123456789',
          },
        })
      );
    }

    const responses = await Promise.all(promises);
    
    // Verify all requests succeeded
    let successCount = 0;
    responses.forEach((response, index) => {
      if (response.ok()) {
        successCount++;
      } else {
        console.log(`Request ${index} failed with status: ${response.status()}`);
      }
      expect(response.ok()).toBeTruthy();
    });

    console.log(`✅ Successfully completed ${successCount}/${responses.length} concurrent requests`);
  });

  test('should fail without authentication', async ({ request }) => {
    const restrictions = {
      maxCaloriesPerSlice: 500,
      mustBeVegetarian: false,
      maxNumberOfToppings: 6,
      minNumberOfToppings: 2
    };

    const response = await request.post(`${BASE_URL}/api/pizza`, {
      data: restrictions,
      headers: {
        'Content-Type': 'application/json',
        // No Authorization header
      },
    });

    // Should get 401 Unauthorized
    expect(response.status()).toBe(401);
    
    const errorBody = await response.json();
    expect(errorBody).toHaveProperty('error');
    expect(errorBody.error).toBe('authentication failed');
    
    console.log(`✅ Correctly rejected unauthenticated request: ${errorBody.error}`);
  });
});
