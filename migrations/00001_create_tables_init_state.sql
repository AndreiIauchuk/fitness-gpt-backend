-- +goose Up
CREATE TABLE fitness.User_Info (
    Id int primary key,
    Name varchar(30),
    Surname varchar(30),
    PhoneNumber varchar(30),
    Address varchar(255)
);

CREATE TABLE fitness.User_Param (
    Id int primary key,
    Weight int,
    Height int,
    SkillLevel int --Пока не понятно, хорошая ли идея хранить уровень тренированности юзера тут и вообще нужно ли
);

CREATE TABLE fitness.Role (
    Id int primary key,
    Name varchar(30)
);

CREATE TABLE fitness.User_Exercise (
  Id int primary key,
  UserId int,
  ExerciseId int
);

CREATE TABLE fitness.Exercise (
  Id int primary key,
  Name varchar(100),
  EquipmentId int,
  RepetitionCount int
);

CREATE TABLE fitness.Equipment (
  Id int primary key,
  Name varchar(100)
);

CREATE TABLE fitness.Exercise_Muscle (
  Id int primary key,
  ExerciseId int,
  MuscleId int
);

CREATE TABLE fitness.Muscle (
  Id int primary key,
  Name varchar(100),
  MuscleGroupId int
);

CREATE TABLE fitness.Muscle_Group (
  Id int primary key,
  Name varchar(100)
);

-- +goose Down
DROP TABLE User_Info;
DROP TABLE User_Param;
DROP TABLE Role;
DROP TABLE User_Exercise;
DROP TABLE Exercise;
DROP TABLE Equipment;
DROP TABLE Exercise_Muscle;
DROP TABLE Muscle;
DROP TABLE Muscle_Group;