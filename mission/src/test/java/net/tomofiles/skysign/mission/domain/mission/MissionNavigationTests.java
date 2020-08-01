package net.tomofiles.skysign.mission.domain.mission;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotEquals;

public class MissionNavigationTests {
    
    /**
     * Mission Navigationの1件同士の比較結果が同一であること。
     */
    @Test
    public void sameTwoMissionNavigationInSingleTest() {
        Navigation navigationA = new Navigation();
        Navigation navigationB = new Navigation();

        navigationA.setTakeoffPointGroundHeight(Height.fromM(0.0));
        navigationA.pushNextWaypoint(
            new GeodesicCoordinates(1.0, 2.0),
            Height.fromM(3.0),
            Speed.fromMS(4.0));

        navigationB.setTakeoffPointGroundHeight(Height.fromM(0.0));
        navigationB.pushNextWaypoint(
            new GeodesicCoordinates(1.0, 2.0),
            Height.fromM(3.0),
            Speed.fromMS(4.0));

        assertEquals(navigationA, navigationB);
    }
    
    /**
     * Mission Navigationの複数件かつ同順序同士の比較結果が同一であること。
     */
    @Test
    public void sameTwoMissionNavigationInSeveralInSameOrderTest() {
        Navigation navigationA = new Navigation();
        Navigation navigationB = new Navigation();

        navigationA.setTakeoffPointGroundHeight(Height.fromM(0.0));
        navigationA.pushNextWaypoint(
            new GeodesicCoordinates(1.0, 2.0),
            Height.fromM(3.0),
            Speed.fromMS(4.0));
        navigationA.pushNextWaypoint(
            new GeodesicCoordinates(11.0, 12.0),
            Height.fromM(13.0),
            Speed.fromMS(14.0));
        navigationA.pushNextWaypoint(
            new GeodesicCoordinates(21.0, 22.0),
            Height.fromM(23.0),
            Speed.fromMS(24.0));

        navigationB.setTakeoffPointGroundHeight(Height.fromM(0.0));
        navigationB.pushNextWaypoint(
            new GeodesicCoordinates(1.0, 2.0),
            Height.fromM(3.0),
            Speed.fromMS(4.0));
        navigationB.pushNextWaypoint(
            new GeodesicCoordinates(11.0, 12.0),
            Height.fromM(13.0),
            Speed.fromMS(14.0));
        navigationB.pushNextWaypoint(
            new GeodesicCoordinates(21.0, 22.0),
            Height.fromM(23.0),
            Speed.fromMS(24.0));

        assertEquals(navigationA, navigationB);
    }
    
    /**
     * Mission Navigationの複数件かつ異なる順序同士の比較結果が同一でないこと。
     */
    @Test
    public void differentTwoMissionNavigationInSeveralInAnotherOrderTest() {
        Navigation navigationA = new Navigation();
        Navigation navigationB = new Navigation();

        navigationA.setTakeoffPointGroundHeight(Height.fromM(0.0));
        navigationA.pushNextWaypoint(
            new GeodesicCoordinates(1.0, 2.0),
            Height.fromM(3.0),
            Speed.fromMS(4.0));
        navigationA.pushNextWaypoint(
            new GeodesicCoordinates(11.0, 12.0),
            Height.fromM(13.0),
            Speed.fromMS(14.0));
        navigationA.pushNextWaypoint(
            new GeodesicCoordinates(21.0, 22.0),
            Height.fromM(23.0),
            Speed.fromMS(24.0));

        navigationB.setTakeoffPointGroundHeight(Height.fromM(0.0));
        navigationB.pushNextWaypoint(
            new GeodesicCoordinates(11.0, 12.0),
            Height.fromM(13.0),
            Speed.fromMS(14.0));
        navigationB.pushNextWaypoint(
            new GeodesicCoordinates(1.0, 2.0),
            Height.fromM(3.0),
            Speed.fromMS(4.0));
        navigationB.pushNextWaypoint(
            new GeodesicCoordinates(21.0, 22.0),
            Height.fromM(23.0),
            Speed.fromMS(24.0));

        assertNotEquals(navigationA, navigationB);
    }
}