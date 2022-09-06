import React from "react";
import { Image, View } from "react-native";
import { TouchableOpacity } from "react-native-gesture-handler";

import certifiedIconSVG from "../../assets/icons/certified.svg";
import { Collection } from "../api/marketplace/v1/marketplace";
import { useAppNavigation } from "../utils/navigation";
import { BrandText } from "./BrandText";
import { SVG } from "./SVG";
import { CardOutline } from "./cards/CardOutline";

export const collectionItemHeight = 266;
export const collectionItemWidth = 196;

export const CollectionView: React.FC<{
  item: Collection;
}> = ({ item }) => {
  const navigation = useAppNavigation();
  return (
    <TouchableOpacity
      onPress={() =>
        navigation.navigate("Collection", { mintAddress: item.mintAddress })
      }
      disabled={!item.mintAddress}
    >
      <CardOutline
        style={{
          paddingTop: 12,
          paddingBottom: 20,
          width: collectionItemWidth,
          height: collectionItemHeight,
        }}
      >
        <Image
          source={{ uri: item.imageUri }}
          style={{
            width: 172,
            height: 172,
            alignSelf: "center",
            borderRadius: 12,
          }}
        />
        <View style={{ marginHorizontal: 12, marginTop: 16 }}>
          <BrandText
            style={{ fontSize: 14 }}
            ellipsizeMode="tail"
            numberOfLines={1}
          >
            {item.collectionName}
          </BrandText>
          <View
            style={{
              flexDirection: "row",
              alignItems: "center",
              marginTop: 8,
            }}
          >
            <BrandText
              style={{ color: "#AEB1FF", fontSize: 14 }}
              ellipsizeMode="tail"
              numberOfLines={1}
            >
              {item.creatorName}
            </BrandText>
            {item.verified && (
              <View style={{ marginLeft: 14 }}>
                <SVG source={certifiedIconSVG} />
              </View>
            )}
          </View>
        </View>
      </CardOutline>
    </TouchableOpacity>
  );
};
