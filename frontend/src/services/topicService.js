const API_BASE_URL = "http://localhost:8080"; 

export async function fetchTopics() {
  try {
    const response = await fetch(`${API_BASE_URL}/topics`); 
    if (!response.ok) {
      throw new Error("Failed to fetch topics");
    }
    const data = await response.json();
    console.log("Fetched topics data:", data);
    return data;
  } catch (error) {
    console.error("Error fetching topics:", error);
    return [];
  }
}
