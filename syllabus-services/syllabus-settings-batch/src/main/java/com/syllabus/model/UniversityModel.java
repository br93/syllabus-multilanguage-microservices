package com.syllabus.model;

import java.time.Instant;

import org.springframework.format.annotation.DateTimeFormat;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.Table;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Entity
@Table(name = "tb_universities")
@NoArgsConstructor
@AllArgsConstructor
@Data
public class UniversityModel {

    @Id
	@GeneratedValue(strategy = GenerationType.IDENTITY)
	private Long id;
	
	@DateTimeFormat(pattern = "yyyy-MM-dd HH:mm:ss")
	private Instant createdAt;
	
	@DateTimeFormat(pattern = "yyyy-MM-dd HH:mm:ss")
	private Instant updatedAt;
	
	@DateTimeFormat(pattern = "yyyy-MM-dd HH:mm:ss")
	private Instant deletedAt;
	
	private String universityId;
	private String universityName;
	private String universityCode;
	
    
}
