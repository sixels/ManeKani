import { Meta, StoryObj } from "@storybook/react";
import { Button } from "./Button";
import { IconShare } from "@tabler/icons-react";

const meta: Meta<typeof Button> = {
	title: "Primitives/Button",
	component: Button,
	tags: ["autodocs"],
	argTypes: {
		icon: {
			options: ["None", "Share"],
			mapping: {
				None: undefined,
				Share: <IconShare strokeWidth={1.25} size={20} />,
			},
		},
		iconOnLeft: {
			defaultValue: "right",
		},
	},
};

export default meta;

type Story = StoryObj<typeof Button>;

export const Primary: Story = {
	args: {
		text: "Primary Button",
		isPrimary: true,
		fillWidth: false,
		size: "md",
	},
};

export const Secondary: Story = {
	args: {
		text: "Secondary Button",
		isPrimary: false,
		fillWidth: false,
		size: "md",
	},
};

export const FullWidth: Story = {
	args: {
		text: "Full Width Button",
		isPrimary: true,
		fillWidth: true,
		size: "md",
	},
};

export const ExtraSmall: Story = {
	args: {
		text: "Small Button",
		isPrimary: false,
		fillWidth: false,
		size: "xs",
		icon: <IconShare size={16} strokeWidth={1.25} />,
		iconOnLeft: true,
	},
};

export const ExtraLarge: Story = {
	args: {
		text: "Large Button",
		isPrimary: true,
		fillWidth: false,
		size: "xl",
	},
};
