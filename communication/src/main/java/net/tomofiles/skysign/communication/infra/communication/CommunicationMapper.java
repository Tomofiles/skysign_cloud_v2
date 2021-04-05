package net.tomofiles.skysign.communication.infra.communication;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface CommunicationMapper {
    @Select("SELECT id FROM communication")
    List<CommunicationRecord> findAll();

    @Select("SELECT id FROM communication WHERE id = #{id} FOR UPDATE")
    CommunicationRecord find(String id);

    @Insert("INSERT INTO communication (id) VALUES (#{id})")
    void create(CommunicationRecord communication);

    @Update("DELETE FROM communication WHERE id = #{id}")
    void delete(String id);
}