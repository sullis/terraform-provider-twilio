terraform {
  required_providers {
    twilio = {
      source  = "twilio/twilio"
      version = ">=0.4.0"
    }
  }
}

provider "twilio" {
  //  username defaults to TWILIO_API_KEY with TWILIO_ACCOUNT_SID as the fallback env var
  //  password  defaults to TWILIO_API_SECRET with TWILIO_AUTH_TOKEN as the fallback env var
  //  account_sid  defaults to TWILIO_SUBACCOUNT_SID with TWILIO_ACCOUNT_SID as the fallback env var
}


resource "twilio_api_accounts_incoming_phone_numbers_v2010" "phone_number" {
  area_code     = "415"
  friendly_name = "terraform phone number"
  sms_url       = "https://demo.twilio.com/welcome/sms/reply"
  voice_url     = "https://demo.twilio.com/docs/voice.xml"
}


resource "twilio_api_accounts_messages_v2010" "message" {
  path_account_sid = "AC123456123456123456123456123456"
  from             = twilio_api_accounts_incoming_phone_numbers_v2010.phone_number.phone_number
  to               = "4444444444"
  body             = "Hello from Twilio Terraform!"
}

resource "twilio_api_accounts_calls_v2010" "call" {
  from  = twilio_api_accounts_incoming_phone_numbers_v2010.phone_number.phone_number
  to    = "4444444444"
  twiml = "<?xml version=\"1.0\" encoding=\"UTF-8\"?><Response><Say>Hello World</Say></Response>"
}

output "phone_numbers" {
  value = twilio_api_accounts_incoming_phone_numbers_v2010.phone_number
}

output "messages" {
  value = twilio_api_accounts_messages_v2010.message
}

output "calls" {
  value = twilio_api_accounts_calls_v2010.call
}

