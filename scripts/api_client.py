#!/usr/bin/env python3

import requests
from typing import Dict, Any, Optional

class APIClient:
    def __init__(self, base_url: str, api_key: str):
        self.base_url = base_url
        self.headers = {
            "Content-Type": "application/json",
            "X-API-Key": api_key
        }

    def make_request(self, method: str, endpoint: str, data: Dict[str, Any] = None) -> Dict[str, Any]:
        url = f"{self.base_url}/api{endpoint}"
        
        try:
            if method == "GET":
                response = requests.get(url, headers=self.headers, verify=False)
            elif method == "POST":
                response = requests.post(url, json=data, headers=self.headers, verify=False)
            else:
                raise ValueError(f"Unsupported HTTP method: {method}")
            
            response.raise_for_status()
            return response.json()
        except requests.exceptions.RequestException as e:
            print(f"Error making request to {url}: {e}")
            if hasattr(e.response, 'text'):
                print(f"Response: {e.response.text}")
            raise

    def create_customer(self, name: str, employer_id: Optional[int] = None) -> Dict[str, Any]:
        data = {"name": name}
        if employer_id is not None:
            data["employer_id"] = employer_id
        return self.make_request("POST", "/customers", data)

    def create_fund(self, name: str) -> Dict[str, Any]:
        data = {"name": name}
        return self.make_request("POST", "/funds", data)

    def get_all_funds(self) -> Dict[str, Any]:
        return self.make_request("GET", "/funds")

    def create_investment(self, client_id: int, fund_id: int, amount: float) -> Dict[str, Any]:
        data = {
            "client_id": client_id,
            "fund_id": fund_id,
            "amount": amount
        }
        return self.make_request("POST", "/investments", data)

    def get_investment(self, investment_id: int) -> Dict[str, Any]:
        return self.make_request("GET", f"/investments/{investment_id}")

    def create_employer(self, name: str) -> Dict[str, Any]:
        data = {"name": name}
        return self.make_request("POST", "/employers", data) 