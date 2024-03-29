package com.syllabus.unmarshal.impl;

import org.springframework.stereotype.Component;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.syllabus.data.StudentResponse;
import com.syllabus.unmarshal.Unmarshal;

@Component
public class StudentUnmarshal implements Unmarshal<StudentResponse> {

    @Override
    public StudentResponse toResponse(Object object) {
        return new ObjectMapper().convertValue(object, new TypeReference<StudentResponse>() {
        });
    }

}
