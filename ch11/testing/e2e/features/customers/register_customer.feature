Feature: Register Customer

  Scenario: Registering a new customer
    Given no customer named "John Smith" exists
    When I register a new customer as "John Smith" and number "555-1212"
    Then I expect the request to succeed
    And expect a customer named "John Smith" to exist
