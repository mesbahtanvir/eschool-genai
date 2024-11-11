import React from "react";
import styled from "styled-components";

// Styled components
const Container = styled.section`
  max-width: 800px;
  margin: auto;
  padding: 20px;
  font-family: Arial, sans-serif;
  color: #333;
  background-color: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
`;

const Title = styled.h1`
  font-size: 2rem;
  color: #333;
  text-align: center;
  margin-bottom: 10px;
`;

const Overview = styled.p`
  font-size: 1.1rem;
  color: #555;
  margin-bottom: 20px;
  line-height: 1.5;
`;

const ModuleTitle = styled.h2`
  font-size: 1.5rem;
  color: #444;
  margin-top: 20px;
  margin-bottom: 10px;
`;

const ModulesWrapper = styled.div`
  margin-top: 10px;
`;

const ModuleContainer = styled.article`
  background-color: #fff;
  border-radius: 8px;
  padding: 15px;
  margin-bottom: 15px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
`;

const ModuleDescription = styled.p`
  font-size: 1rem;
  color: #555;
  margin-top: 10px;
  line-height: 1.5;
`;

// Main CourseBlueprint component
export const CourseBlueprint = ({ data }) => {
  if (!data?.course?.course_blueprint) return <p>Loading course data...</p>;

  const { title, description, modules } = data.course.course_blueprint;

  return (
    <Container>
      <Title>{title}</Title>
      <Overview>{description}</Overview>
      <ModuleTitle>Course Modules</ModuleTitle>
      <ModulesWrapper>
        {modules.map((module, i) => (
          <ModuleContainer key={i}>
            <ModuleTitle>{module.title}</ModuleTitle>
            {module.description && (
              <ModuleDescription>{module.description}</ModuleDescription>
            )}
          </ModuleContainer>
        ))}
      </ModulesWrapper>
    </Container>
  );
};

export default CourseBlueprint;
