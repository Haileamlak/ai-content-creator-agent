# AI Content Creator Agent

**Description**  
This project is an AI-driven content creation and social media management system designed to automate tasks like writing posts, managing interactions, and performing research. The system integrates with social media platforms to post, respond to comments, and analyze engagement trends, all while leveraging AI for personalized, high-quality content generation.

## Features
- Automated content creation using **Gemini API**
- Social media management across **Twitter and Facebook**
- Research module using **News API** for trending content
- Sentiment analysis and performance tracking for future content improvements
- Scheduled posting and interaction using **Google Cloud Scheduler**

## Technologies
- **Backend:** Golang
- **APIs:** Gemini, Twitter, Facebook, News API
- **Database:** Google Cloud Firestore
- **Scheduling:** Google Cloud Scheduler
- **Cloud Platform:** Google Cloud Project (GCP)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/haileamlak/ai-content-creator-agent.git
   cd ai-content-creator-agent
   ```

2. Install dependencies for Golang services:
   ```bash
   go mod download
   ```

3. Set up environment variables for Google Cloud and external APIs:
   - Add your service account key JSON file for Firestore and Scheduler.
   - Set API keys for Gemini, News API, and social media platforms in your `.env` file.

4. Run the backend server:
   ```bash
   go run main.go
   ```

## Usage

- After setup, the system will automatically generate and post content, interact with users, and analyze performance.
- Posts can be scheduled via the Cloud Scheduler or manually triggered via the API.

## Contributing
Feel free to contribute by opening a pull request or submitting an issue.

## License
This project is licensed is not licensed under any type of License.
