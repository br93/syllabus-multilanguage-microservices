package com.syllabus.client.settings.response;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@NoArgsConstructor
@AllArgsConstructor
@Data
@JsonIgnoreProperties(ignoreUnknown = true)
public class PreRequisiteCoursesResponse {

    @JsonProperty("course_id")
    private String courseId;

    @JsonProperty("course_code")
    private String courseCode;

    @JsonProperty("course_name")
    private String courseName;

    @JsonProperty("workload")
    private Short workload;

    @JsonProperty("pre_requisite_courses")
    private List<PreRequisiteResponse> preRequisites;    
}
