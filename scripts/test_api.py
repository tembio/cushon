#!/usr/bin/env python3

import json
import urllib3
from api_client import APIClient

# Suppress only the single warning from urllib3 needed.
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

def main():
    # Initialize API client
    client = APIClient(
        base_url="https://localhost:8443",
        api_key="test-api-key"
    )

    # Create an employer
    print("\nCreating employer...")
    employer = client.create_employer("Tech Corp")
    print(f"Created employer: {json.dumps(employer, indent=2)}")
    employer_id = employer["id"]

    # Create two customers - one retail and one employed
    print("\nCreating retail customer...")
    retail_customer = client.create_customer("John Doe")
    print(f"Created retail customer: {json.dumps(retail_customer, indent=2)}")
    retail_customer_id = retail_customer["id"]

    print("\nCreating employed customer...")
    employed_customer = client.create_customer("Jane Smith", employer_id)
    print(f"Created employed customer: {json.dumps(employed_customer, indent=2)}")
    employed_customer_id = employed_customer["id"]

    # Create multiple funds
    print("\nCreating funds...")
    
    fund1 = client.create_fund("Fund1")
    print(f"Created fund: {json.dumps(fund1, indent=2)}")
    fund1_id = fund1["id"]

    fund2 = client.create_fund("Fund2")
    print(f"Created fund: {json.dumps(fund2, indent=2)}")
    fund2_id = fund2["id"]

    fund3 = client.create_fund("Fund3")
    print(f"Created fund: {json.dumps(fund3, indent=2)}")
    fund3_id = fund3["id"]

    # Get all funds to verify
    print("\nGetting all funds...")
    all_funds = client.get_all_funds()
    print(f"All funds: {json.dumps(all_funds, indent=2)}")

    # Create investments for both customers
    print("\nCreating investments...")
    
    # Investments for retail customer
    retail_investment1 = client.create_investment(
        client_id=retail_customer_id,
        fund_id=fund1_id,
        amount=2000.0
    )
    print(f"Created retail customer investment 1: {json.dumps(retail_investment1, indent=2)}")

    retail_investment2 = client.create_investment(
        client_id=retail_customer_id,
        fund_id=fund3_id,
        amount=1500.0
    )
    print(f"Created retail customer investment 2: {json.dumps(retail_investment2, indent=2)}")

    # Investment for employed customer
    employed_investment = client.create_investment(
        client_id=employed_customer_id,
        fund_id=fund2_id,
        amount=3000.0
    )
    print(f"Created employed customer investment: {json.dumps(employed_investment, indent=2)}")

    # Get individual investments
    print("\nGetting individual investments...")
    retrieved_retail_investment1 = client.get_investment(retail_investment1["id"])
    print(f"Retrieved retail customer investment 1: {json.dumps(retrieved_retail_investment1, indent=2)}")

    retrieved_retail_investment2 = client.get_investment(retail_investment2["id"])
    print(f"Retrieved retail customer investment 2: {json.dumps(retrieved_retail_investment2, indent=2)}")

    retrieved_employed_investment = client.get_investment(employed_investment["id"])
    print(f"Retrieved employed customer investment: {json.dumps(retrieved_employed_investment, indent=2)}")

    # Get all investments for each customer
    print("\nGetting all investments for retail customer...")
    retail_investments = client.get_investments_by_client(retail_customer_id)
    print(f"All investments for retail customer: {json.dumps(retail_investments, indent=2)}")

    print("\nGetting all investments for employed customer...")
    employed_investments = client.get_investments_by_client(employed_customer_id)
    print(f"All investments for employed customer: {json.dumps(employed_investments, indent=2)}")

if __name__ == "__main__":
    main() 