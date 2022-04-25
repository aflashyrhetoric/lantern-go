ALTER TABLE people
  ADD relationship_to_user text;

ALTER TABLE people
  ADD relationship_to_user_through_person_id int;