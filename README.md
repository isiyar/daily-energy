# Daily Energy Telegram Mini-App

## Overview
The Daily Energy Telegram Mini-App is a user-friendly tool designed to help individuals track their daily activity, calorie intake, and energy levels. It offers an intuitive interface for setting personal details, monitoring progress, and receiving personalized feedback from an AI assistant.

## Functionality

### Registration
The initial setup process guides users through entering personal details to create a profile. Upon successful registration, the AI generates a personalized 7-day plan for food and activity based on the provided data.  
<p align="center">
  <img src="misc/img/greeting.png" height=300px alt="Welcome Screen">  
  <img src="misc/img/input_name.png" height=300px alt="Name Input">  
  <img src="misc/img/choose_gender.png" height=300px alt="Gender Selection">  
  <img src="misc/img/choose_aim.png" height=300px alt="Goal Selection">  
  <img src="misc/img/choose_weight.png" height=300px alt="Weight Input">  
  <img src="misc/img/choose_height.png" height=300px alt="Height Input">  
  <img src="misc/img/choose_day_of_birth.png" height=300px alt="Date of Birth Selection">  
  <img src="misc/img/choose_physical_activity.png" height=300px alt="Activity Level Selection">  
</p>

### Interactive Chat with the AI Assistant Rafik (Mascot)
Users can engage in an interactive chat with Rafik, the AI mascot, for personalized advice and support.  
<p align="center">
  <img src="misc/img/ai_chat_start.png" height=300px alt="Chat Start">  
  <img src="misc/img/ai_chat.png" height=300px alt="Chat Interface">  
</p>

### View and Add Today's Food and Activities
Track and log daily food intake and physical activities. When entering a food name, the AI calculates the average calorie content of a typical serving.  
<p align="center">
  <img src="misc/img/view_today_activity.png" height=300px alt="View Today's Activity">  
  <img src="misc/img/view_today_food.png" height=300px alt="View Today's Food">  
</p>
<p align="center">
  <img src="misc/img/add_today_activity.png" height=300px alt="Add Today's Activity">  
  <img src="misc/img/add_today_food.png" height=300px alt="Add Today's Food">  
</p>

### View History and Plan for Food and Activities
Review historical data and access the 7-day plan generated during registration for both food and activities.  
<p align="center">
  <img src="misc/img/view_history_activity.png" height=300px alt="View Activity History">  
  <img src="misc/img/view_history_food.png" height=300px alt="View Food History">
</p>
<p align="center">
  <img src="misc/img/view_activity_plan.png" height=300px alt="View Activity Plan">  
  <img src="misc/img/view_food_plan.png" height=300px alt="View Food Plan">  
</p>

### Calendar
Select a specific day to view its history or plan details.  
<p align="center">
  <img src="misc/img/calendar.png" height=300px alt="Calendar">  
</p>

### Profile
View personal progress dynamics and edit profile details as needed.  
<p align="center">
  <img src="misc/img/profile.png" height=300px alt="Profile Overview">  
  <img src="misc/img/edit_profile.png" height=300px alt="Edit Profile">  
</p>

## ER-diagram

<img src="misc/img/er-diagram.svg" height=500px alt="Edit Profile">

## Techologies

1. Golang
2. Typescript
3. gin
4. gorm
5. PostgreSQL
6. React
7. React Query
8. TailwindCSS
9. nginx
10. docker compose

As AI model we use openrouter/cypher-alpha:free.
