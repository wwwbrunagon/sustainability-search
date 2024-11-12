// Topics.js
import React, { useEffect, useState } from "react";
import { fetchTopics } from "../services/topicService";
import "./Topics.css"; // Import the CSS file

const Topics = () => {
  const [topics, setTopics] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    async function loadTopics() {
      try {
        const data = await fetchTopics();
        setTopics(Array.isArray(data) ? data : []);
      } catch (err) {
        setError("Failed to load topics");
      }
    }
    loadTopics();
  }, []);

  return (
    <div>
      <h1>Topics:</h1>
      {error ? (
        <p className="error-message">{error}</p>
      ) : (
        <div className="topic-list">
          {topics.length > 0 ? (
            topics.map((topic) => (
              <div key={topic.id} className="topic-card">
                <h2>{topic.name}</h2>
                <p>{topic.description}</p>
              </div>
            ))
          ) : (
            <p>No topics available.</p>
          )}
        </div>
      )}
    </div>
  );
};

export default Topics;
