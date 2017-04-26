package model

import (
	"testing"
	"gopkg.in/yaml.v2"
	"gopkg.in/stretchr/testify.v1/assert"
)

var railgun150mmII string = `3074:
    basePrice: 183136.0
    capacity: 0.1
    description:
        de: 'Dieses ist eine Standardlangstreckenrailgun, die für Fregatten entwickelt
            wurde. Railguns verwenden Magnetschienen, um solide Materiebrocken mit
            Überschallgeschwindigkeit zu verschießen. Die Präzisionsreichweite von
            Railguns ist sehr gut, aber aufgrund technischer Einschränkungen kann
            keine bordeigene Zielsuche benutzt werden. Das führt zu einem drastischen
            Abfall der Genauigkeit bei extremen Entfernungen




            Benötigt Hybridmunition (regulär oder fortschrittlich): Antimatter, Iridium,
            Iron, Lead, Plutonium, Thorium, Tungsten, Uranium, Javelin, Spike.'
        en: "This is a standard long-range railgun designed for frigates. Railguns
            use magnetic rails to fire solid chunks of matter at hypersonic speed.
            The accurate range of railguns is very good, but due to technical limitations
            it cannot use onboard guidance. This results in a fairly rapid drop in
            accuracy at extreme ranges. \r\n\r\nRequires either regular or advanced
            hybrid charge types: Antimatter, Iridium, Iron, Lead, Plutonium, Thorium,
            Tungsten, Uranium, Javelin, Spike."
        fr: "Canon à rail standard de longue portée destiné aux frégates. Les canons
            à rail font appel à des rails magnétiques pour envoyer de gros morceaux
            de matières à une vitesse hypersonique. La portée des canons à rail est
            très précise, mais des limites techniques empêchent l'utilisation de systèmes
            de guidage embarqués. Cela se traduit par une diminution notable de la
            précision à très longue portée. \n\nRequiert les types de charges hybrides
            normales ou avancées suivants : antimatière, iridium, fer, plomb, plutonium,
            thorium, tungstène, uranium, javelin, pieux."
        ja: 'これはフリゲート用の標準的な長射程レールガンである。レールガンは電磁誘導レールを使って固体弾を超音速で発射する兵器である。レールガンの命中精度は非常に高いが、技術上の制約で、弾に誘導装置を搭載できない。そのため超長距離射程では命中精度が急激に低下する。




            装填できる標準型/高性能ハイブリッド弾のタイプ:アンチマター弾、イリジウム弾、アイアン弾、鉛弾、プルトニウム弾、トリウム弾、タングステン弾、ウラン弾、ジャベリン弾、スパイク弾。'
        ru: "Стандартный рейлган большой дальности действия. Предназначен для оснащения
            фрегатов. Рейлган — это электромагнитный ускоритель, «рельсы» которого
            способны разогнать снаряд до гиперзвуковой скорости. Из рейлгана можно
            вести прицельный огонь на очень большие дистанции, однако по мере удаления
            цели точность стрельбы по ней будет заметно снижаться — увы, снаряды рейлганов
            нельзя оборудовать системами наведения. \n\n\n\nСтрельба из рейлгана ведётся
            стандартными и усовершенствованными гибридными боеприпасами следующих
            типов: Antimatter, Iridium, Iron, Lead, Plutonium, Thorium, Tungsten,
            Uranium, Javelin, Spike."
        zh: 这是为护卫舰设计的标准远程磁轨炮。使用电磁轨道将实心的弹丸以超高音速发射出去。磁轨炮的精确射程很长，但由于技术限制，不能进行舰载制导。这导致了在极限射程附近命中率会骤减。使用的普通或高级混合弹药类型：反物质弹药、铱质弹药、铁质弹药、铅质弹药、锰质弹药、钍质弹药、钨质弹药、钼质弹药、“标枪”、“钉刺”。
    graphicID: 11259
    groupID: 74
    iconID: 349
    marketGroupID: 564
    mass: 500.0
    name:
        de: 150mm Railgun II
        en: 150mm Railgun II
        fr: Canon à rail 150mm II
        ja: 150mmレールガンII
        ru: 150mm Railgun II
        zh: 150mm磁轨炮 II
    portionSize: 1
    published: true
    radius: 10.0
    sofFactionName: laidai
    volume: 5.0`

func TestLoad150mmRailgunII(t *testing.T) {
	assert := assert.New(t)

	theTypes := make(map[int32]*Type)
	err := yaml.Unmarshal([]byte(railgun150mmII), theTypes)

	if err != nil {
		t.Error("Received an error when one wasn't expected")
	}

	theType := theTypes[3074]

	assert.NotNil(theType)
	assert.True(theType.BasePrice >= 183136.0 && theType.BasePrice <= 183136.1, "Base price mismatch")
	assert.True(theType.Capacity >= 0.1 && theType.Capacity < 0.2, "Capacity mismatch")
	assert.Equal("This is a standard long-range railgun designed for frigates. Railguns use magnetic rails to fire solid chunks of matter at hypersonic speed. The accurate range of railguns is very good, but due to technical limitations it cannot use onboard guidance. This results in a fairly rapid drop in accuracy at extreme ranges. \r\n\r\nRequires either regular or advanced hybrid charge types: Antimatter, Iridium, Iron, Lead, Plutonium, Thorium, Tungsten, Uranium, Javelin, Spike.", theType.Description["en"], "EN Description Mismatch")
	assert.Equal(int32(11259), theType.GraphicId, "Graphic ID Mismatch")
	assert.Equal(349, theType.IconId, "Icon ID Mismatch")
	assert.Equal(564, theType.MarketGroupId, "Market Group ID Mismatch")
	assert.True(theType.Mass >= 500.0 && theType.Mass < 500.1, "Mass Mismatch")
	assert.Equal("150mm Railgun II", theType.Name["en"], "EN Name Mismatch")
	assert.Equal(1, theType.PortionSize, "Portion Size Mismatch")
	assert.Equal(true, theType.Published, "Published Mismatch")
	assert.True(theType.Radius >= 10.0 && theType.Radius < 10.1, "Radius Mismatch")
	assert.Equal("laidai", theType.FactionName, "Faction Name Mismatch")
	assert.True(theType.Volume >= 5.0 && theType.Volume < 5.1, "Volume Mismatch")
}
