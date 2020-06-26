package net.tomofiles.skysign.communication.infra.communication;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import net.tomofiles.skysign.communication.domain.common.Version;
import net.tomofiles.skysign.communication.domain.communication.Communication;
import net.tomofiles.skysign.communication.domain.communication.CommunicationFactory;
import net.tomofiles.skysign.communication.domain.communication.CommunicationId;
import net.tomofiles.skysign.communication.domain.communication.CommunicationRepository;

@Component
public class CommunicationRepositoryImpl implements CommunicationRepository {

    @Autowired
    private CommunicationMapper communicationMapper;

    @Override
    public void save(Communication communication) {
        boolean isCreate = false;

        CommunicationRecord record = this.communicationMapper.find(communication.getId().getId());

        if (record == null) {
            record = new CommunicationRecord();
            record.setId(communication.getId().getId());
            isCreate = true;
        }

        record.setMissionId(communication.getMissionId() == null ? null : communication.getMissionId().getId());
        record.setVersion(communication.getVersion().getVersion());

        if (isCreate) {
            this.communicationMapper.create(record);
        } else {
            this.communicationMapper.update(record);
        }
    }

    @Override
    public void remove(CommunicationId id, Version version) {
        DeleteCondition condition = new DeleteCondition();
        condition.setId(id.getId());
        condition.setVersion(version.getVersion());
        this.communicationMapper.delete(condition);
    }

    @Override
    public Communication getById(CommunicationId id) {
        CommunicationRecord record = this.communicationMapper.find(id.getId());

        if (record == null) {
            return null;
        }

        return CommunicationFactory.rebuild(
            id, 
            record.getMissionId(), 
            record.getVersion());
    }
}